package main

import (
	"fmt"
	"testing"
)

type Any interface {
}

type EvalFunc func(initState Any) (Any, Any)

func TestLazyEval(t *testing.T) {
	evalFunc := func(initState Any) (result Any, nextState Any) {
		result = initState
		nextState = initState.(int) + 2
		return
	}
	even := buildLazyIntEvaluator(evalFunc, 0)
	for i := 0; i < 10; i++ {
		fmt.Println(even())
	}
}

func buildLazyEvaluator(evalFunc EvalFunc, initState Any) chan Any {
	outChan := make(chan Any)
	loopFunc := func() {
		activeState := initState
		var result Any
		for {
			result, activeState = evalFunc(activeState)
			outChan <- result
		}
	}
	go loopFunc()
	return outChan
}

func buildLazyIntEvaluator(evalFunc EvalFunc, initState int) func() int {
	outChan := buildLazyEvaluator(evalFunc, initState)
	return func() int {
		return (<-outChan).(int)
	}
}
