package main
import (
	"fmt"
	"strings"
	"strconv"
	"log"
)

func main(){
	
	var s string
	
	fmt.Println("enter fractions for multiplication in string form as 6/5*4/10")
	fmt.Scanf("%s",&s)
	
	fractions:=strings.Split(s,"*") 
	
	output:=MultiplyFractions(fractions[0],fractions[1])
	fmt.Println("*****MUltiplication Result**********",output)
	
}
func MultiplyFractions(fractionone string,fractiontwo string)string{
	chcomplex:=make(chan string,2)
	chsimple:=make(chan string,2)
	
	slicesimplifiedfract:=make([]string,0)
	fractions:=make([]string,0)
	fractions=append(fractions,fractionone)
	fractions=append(fractions,fractiontwo)
	
	for i:=0;i<len(fractions);i++{
		fmt.Println("********* Actual Fractions ************",fractions[i])
		chcomplex<-fractions[i]
		go reduce(chcomplex,chsimple)//simplifies each fraction by running concurrently for each fraction
		slicesimplifiedfract=append(slicesimplifiedfract,<-chsimple)
	}
	
	output:=multiply(slicesimplifiedfract)
	fmt.Println("*****MUltiplication Result 1 COMPLEX**********",output)
	chcomplex<-output
	go reduce(chcomplex,chsimple)
	close(chcomplex)
	output1:=<-chsimple
	fmt.Println("*****MUltiplication Result 2**SIMPLEX********",output1)
	return output1
}
func reduce(chcomplex <-chan string,chsimple chan<- string){
	var greatest int
	var least int
	var gcd int
	complexfraction:=<-chcomplex
	spfract:=strings.Split(complexfraction,"/")
	numerator,err:=strconv.Atoi(spfract[0])
	if err!=nil{
		fmt.Println(err)
	}
	denominator,err:=strconv.Atoi(spfract[1])
	if err!=nil{
		fmt.Println(err)
	}
	if numerator==0{
		log.Println("*** Final Output will be zero so not simplying pritnting as actual output **")
		gcd=denominator
	}else if denominator==0{
		log.Println("****Denominator is zero so final result should be infinite but printed as actual fractions not simplifying*******")
		gcd=numerator
	}else if numerator!=0 && denominator!=0{
		if numerator>denominator{
			greatest=numerator
			least=denominator
		}else{
			greatest=denominator
			least=numerator
		}
		gcd= findgcd(greatest,least)
		numerator=numerator/gcd
		denominator=denominator/gcd
	}
	
	
	
	simplifiedfraction:=strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator)
	chsimple<-simplifiedfraction
	
}
	func findgcd(dividend int,divisor int)int{
		var gcd int
		
		remainder:=dividend%divisor
		
		if remainder==0{
		gcd=divisor
		}else{
		return findgcd(divisor,remainder)
		
	}
		return gcd
	}
	func multiply(slicesimplifiedfract []string)string{
		c1:=make(chan int)
		c2:=make(chan int)
		numerators:=make([]int,0)
		denominators:=make([]int,0)
		 multiplicationofnum:=1
		 multiplicationofdeno :=1
		
		for i:=range slicesimplifiedfract{
			fmt.Println("***********simplifiedfractions**********",slicesimplifiedfract[i])
			fract:=strings.Split(slicesimplifiedfract[i],"/")
			
			numer,err:=strconv.Atoi(fract[0])
			if err!=nil{
			fmt.Println(err)
			}
		numerators=append(numerators,numer)
		denom,err:=strconv.Atoi(fract[1])
		if err!=nil{
		fmt.Println(err)
		}
		denominators=append(denominators,denom)
		}
		
		
		go func(){
			
			c1<-multiplication(numerators)
		}()
		go func(){
			
			c2<-multiplication(denominators)
		}()
		for i:=0;i<2;i++{
			select{
			case multiplicationofnum=<-c1:
				fmt.Println("multiplication of numerators",multiplicationofnum)
			case multiplicationofdeno=<-c2:
				fmt.Println("multiplication of denominators",multiplicationofdeno)
			}
		}
		
		if multiplicationofnum==0{
			fmt.Println("******  Final RESULT IS ZERO *********")
		}else if multiplicationofdeno==0{
			fmt.Println("******** FINAL RESULT IS INFINITE ****")
		}
		s:=strconv.Itoa(multiplicationofnum)+"/"+strconv.Itoa(multiplicationofdeno)
		
		return s
	}
	func multiplication(operands []int)int{
		result:=1
		for i:=range operands{
		result*=operands[i]
		}
		return result
	}