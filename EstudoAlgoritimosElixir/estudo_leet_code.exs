defmodule PesquisaLinear do
  #exercicio 1
  def primeiro_o(lista, target), do: primeiro_o(lista,target,0)
  def primeiro_o(_lista, _target, {:ok, cnt}), do: cnt
  def primeiro_o([], _target, _cnt), do: :nil
  def primeiro_o([head|tail], target, cnt) do
    primeiro_o(tail, target, cond do
    target == head -> {:ok, cnt+1}
    true -> cnt+1 
    end)
  end
  #exercicio 2
  def contar(lista, elemen), do: contar(lista,elemen, 0)
  def contar([], _elemen, cnt), do: cnt
  def contar([head | tail], elemen, cnt) do
    contar(tail, elemen, cond do
    elemen == head -> cnt+1
    true -> cnt
    end)
  end
  #exercicio 3
  def existe?([],_target), do: :false
  def existe?([head | tail], target) do
    cond do
    head == target -> :true
    true -> existe?(tail, target)
    end
  end
  #exercicio 4
  def index(lista, n), do: index(lista, n, 0)
  def index([], _n, _cnt), do: :nil
  def index([head | tail], n, cnt) do
    if head == n do
    cnt
    else
      new_cnt = cnt+1
      index(tail, n, new_cnt)
    end
  end
  #exercicio 5
  def acha_min([], n), do: n
  def acha_min(lista), do: acha_min(lista, hd(lista))
  def acha_min([head | tail], n) do
    new_n = if head < n, do: head, else: n
    acha_min(tail, new_n)
  end
end

#Novo Conceito

defmodule PesquisaBinaria do
  def
end
