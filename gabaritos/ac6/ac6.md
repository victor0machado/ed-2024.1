PROGRAMA mesclaListas(a, b, m, n)
	i := 0
	j := 0
	Criar vetor v, com m + n elementos
	x := 0 // representa o tamanho de v

	Enquanto i < m e j < n, faça:
		Se a[i] < b[j], então:
			v[x] := a[i]
			i := i + 1
		Senão:
			v[x] := b[j]
			j := j + 1
		x := x + 1

	Enquanto i < m, faça:
		v[x] := a[i]
		x := x + 1
		i := i + 1

	Enquanto j < n, faça:
		v[x] := b[j]
		x := x + 1
		j := j + 1
