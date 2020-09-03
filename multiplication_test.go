package main

import "testing"

type testpair struct{
	fractions []string
	result string

}
var tests=[]testpair{
	{[]string{"1/2","1/3"},"1/6"},
	{[]string{"6/10","4/10"},"6/25"},
	{[]string{"0/6","4/10"},"0/30"},
	{[]string{"4/48","9/54"},"1/72"},
	{[]string{"5/55","8/64"},"1/88"},
	{[]string{"7/28","6/54"},"1/36"},
	{[]string{"8/46","80/100"},"16/115"},
	{[]string{"4/89","78/100"},"78/2225"},
	{[]string{"80/100","68/700"},"68/875"},
	{[]string{"87/115","45/165"},"261/1265"},
	{[]string{"87/0","45/165"},"261/0"},	

}
func TestMultiplyFractions(t *testing.T){
	for _,pair:=range tests{
		v:=MultiplyFractions(pair.fractions[0],pair.fractions[1])
		if v!=pair.result{
			t.Error(
				"for",pair.fractions,
				"expected",pair.result,
				"got",v,
			)
		}
	}

}