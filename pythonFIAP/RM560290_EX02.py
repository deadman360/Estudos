# pede o valor do carro

preco = input("Digite o valor do carro: ")
preco = float(preco)
# define variaveis de parcela e juros
parcelas = 0
juros = 0.0

# define caso a vista
if parcelas == 0:
    print(f"O preço final à vista com desconto de 20% é de R${(preco * 0.80):.2f}")
    parcelas += 6
    juros += 0.03
while parcelas <= 60:
    if parcelas == 6:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 12:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 18:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 24:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 30:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 36:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 42:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 48:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 54:
            valor_parcela = preco * (1 + juros) / parcelas
    elif parcelas == 60:
            valor_parcela = preco * (1 + juros) / parcelas

    total = preco * (1 + juros)
    print(
        f"O Preço final parcelado em X{parcelas} é de R${total:.2f}, com parcelas de R${valor_parcela:.2f}"
    )
    parcelas += 6
    juros += 0.03

    # formatei os numeros por estetica, sei que nao aprendemos isso ainda.
