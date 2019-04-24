<h1>Arrays e Slices em Go</h1>
<p>Em Go arrays tem tamanhao fixo, uma vez definido ele não pode ser alterado.
</p>
<p>var numeros [10]int  === declarando arrays de 10 elementos</p>
<p>Slices são arrays dinamicos que podem aumentar ou diminuir de tamanho conforme for necessário</p>
<p>Internamente um slice utiliza um arrya paa armazenar os seus dados. Um slice possui 3 atributos importantes: ponteiro para o array,
 tamanho atual do slice e capacidade do array.</p>
 <p>meuSlice := make([]int, 3, 5 ) == neste exemplo é criado um slice de tamanho 3 e capacidade 5.</p>


