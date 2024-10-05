#printa escolhas na tela

print("Escolha o tipo de investimento:\n 1. CBD\n 2. LCI\n 3. LCA")
selecionado = input("Digite o tipo de investimento(1,2 ou 3):")
resgate = float(input("Digite o valor a ser resgatado:"))
dias = int(input("Digite o número de dias que permaneceu investido:"))


#define if composto para calculo de aliquota
if selecionado == "1":
    if dias <= 180:
        aliquota = 0.225
    elif dias > 180 and dias <= 360:
        aliquota = 0.20
    elif dias > 360 and dias <= 720:
        aliquota = 0.175
    elif dias > 720:
        aliquota = 0.15
    print(f"O valor do imposto de renda a ser pago é R${resgate * aliquota:.2f}")
else:
    print("Investimento isento de imposto de renda")
