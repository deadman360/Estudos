#pede o valor da dívida

divida = input("Digite o valor da dívida: ")
divida = float(divida)

#estrutura de repetição para calcular o valor de cada parcela

parcelas = 0
juros = 0.0
if parcelas == 0:
    print(f"Total: R${divida:.2f} ,Juros: R${divida * juros:.2f} Numero de parcelas: 1 ,Valor da parcela: R${divida:.2f}")
    parcelas += 3
    

while parcelas <= 12:
    if parcelas == 3:
        juros = 0.10
        valor_parcela = divida * (1 + juros) / parcelas
    elif parcelas == 6:
        valor_parcela = divida * (1 + juros) / parcelas
    elif parcelas == 9:
        valor_parcela = divida * (1 + juros) / parcelas
    elif parcelas == 12:
        valor_parcela = divida * (1 + juros) / parcelas
        
        

    total = divida * (1 + juros)
    print(f"Total: R${total:.2f} ,Juros: R${divida * juros:.2f} Numero de parcelas: {parcelas} ,Valor da parcela: R${valor_parcela:.2f}")
    parcelas += 3
    juros += 0.05
    

    #formatei os numeros por estetica, sei que nao aprendemos isso ainda.
