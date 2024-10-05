#Pede numero de colaboradores e passa o valor para a variavel

colaboradores = input("Informe o número de colaboradores: ")
#Definindo variaveis para cada dia da semana
s = 0
t = 0
qa = 0
qi = 0
sx = 0
#Começa for loop, o contador de cada dia da semana seré representado pela sua letra inicial
for x in range(int(colaboradores)):
    preferencia = input("Informe o dia da sua preferência(Segunda-feira, Terça-feira, Quarta-feira, Quinta-feira, Sexta-feira: ")
    #Tratamento de dados par garantir que o input possa ser digitado de qualquer forma
    preferencia= preferencia.lower()
    preferencia = preferencia.replace("-", "")
    preferencia = preferencia.replace("feira", "")
    preferencia = preferencia.replace("ç", "c")
    preferencia = preferencia.replace(" ", "")

    # aqui usei "print(preferencia)" a fim de testar se os dados estavam sendo tratados corretamente, comentei ele depois

    if preferencia == "segunda":
        s += 1
    elif preferencia == "terca" or preferencia == "terca-feira":
        t += 1
    elif preferencia == "quarta":
        qa += 1
    elif preferencia == "quinta":
        qi += 1
    elif preferencia == "sexta":
        sx += 1
    else:
        print("Dia da semana inválido, seu voto nâo será computado")

#Define o resultado da votação

if s > t and s > qa and s > qi and s > sx:
    print("O dia escolhido para os colaboradores foi: Segunda-feira")
elif t > s and t > qa and t > qi and t > sx:
    print("O dia escolhido para os colaboradores foi: Terça-feira")
elif qa > s and qa > t and qa > qi and qa > sx:
    print("O dia escolhido para os colaboradores foi: Quarta-feira")
elif qi > s and qi > t and qi > qa and qi > sx:
    print("O dia escolhido para os colaboradores foi: Quinta-feira")
elif sx > s and sx > t and sx > qa and sx > qi:
    print("O dia escolhido para os colaboradores foi: Sexta-feira")
else:
    print("Houve um empate, entre em contato com a administração para realizar a votação novamente")



