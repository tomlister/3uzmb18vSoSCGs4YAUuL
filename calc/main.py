print("Simple Calculator v1.0")
num1 = int(input("[NUM1]: "))
op = ""
while True:
	opinp = input("[Operator](+ - * /): ")
	if opinp in "+-*/":
		op = opinp
		break
	else:
		print("[Operator] -> Error Try Again")
num2 = int(input("[NUM2]: "))
result = 0
if op == "+":
	result = num1+num2
elif op == "-":
	result = num1-num2
if op == "*":
	result = num1*num2
if op == "/":
	result = num1/num2
print("Result: "+str(result))