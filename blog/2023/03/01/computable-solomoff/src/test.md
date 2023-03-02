\section{An distracting epicycle: dealing with Turing machines that take too long or do not halt.}

When thinking about Turing machines, one might consider one particular model, e.g., valid C programs. But in that case, it is easy to consider programs that do not halt, like:\footnote{Readers might find it amusing to run gcc loop.c and check it}:

\begin{verbatim}
void main(){
  while(1){
    ;
  }
}
\end{verbatim}

But then in step 2 of our scheme, we never get to see what bits this program outputs, because it never outputs any bits.

This can easily be fixed as follows:

\begin{enumerate}
	\item Start with a finite set of live turing machines, $\{T_0, ..., T_n\}$, a compute budget $s$ and an empty cache of programs which take too long $C =\{\}$
	\item Run each $T_i$ for $s$ seconds. 
	\item If none of them predict your trail bits, $(B_0, ..., B_m)$, within the given compute budget: 
		\begin{enumerate}
		   \item eliminate the ones who make incorrect predictions.
			 \item keep track of programs which didn't output anything in the cache.
			 \item attempt to compute the first $m$ steps of Turing machine $T_{n+1}$ with $s$ seconds of compute. If it makes correct predictions, keep it in the set of live machines, otherwise move it to the cache.
			 \item increase the compute budget to $s + 1$ and run each machine in the cache for one additional second
			 \item Repeat step 2 until you have one program which has predicted past bits within your compute budget. Eventually this program must exist, since the Turing machine which is producing your trail of bits is by construction computable and non-halting.\footnote{Or at least, it hasn't halted before producing the number of bits that you have seen so far.}
			\item 
		\end{enumerate}
	\item Observe the next bit, purge the machines from your set which don't predict it. If none predict it, GOTO 2.
