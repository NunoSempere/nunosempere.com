A computable version of Solomoff induction
==========================================

Thinking about [Just-in-time Bayesianism](https://nunosempere.com/blog/2023/02/04/just-in-time-bayesianism/) a bit more,  here is a computable approximation to Solomonoff Induction, which converges to the Turing machine generating your trail of bits in finite time. <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script> <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script> <!-- Note: to correctly render this math, compile this markdown with /usr/bin/markdown -f fencedcode -f ext -f footnote -f latex $1 where /usr/bin/markdown is the discount markdown binary https://github.com/Orc/discount http://www.pell.portland.or.us/~orc/Code/discount/ --> 

### The key idea: arrive at the correct hypothesis in finite time

1. Start with a finite set of Turing machines, \(\{T_0, ..., T_n\}\)
2. If none of the \(T_i\) predict your trail bits, \((B_0, ..., B_m)\), compute the first \(m\) steps of Turing machine \(T_{n+1}\). If \(T_{n+1}\) doesn't predict them either, go to \(T_{n+2}\), and so on[^1]
3. Observe the next bit, purge the machines from your set which don't predict it. If none predict it, go to 2.

Then in finite time, you will arrive at a set which only contains the simplest TM which describes the process generating your train of bits. Proof:

- Let \(T_j\) be the simplest machine which describes the process generating your trail of bits. Then by virtue of it being such, all TM machines \(\{T_1,...,T_{j-1}\}\) are not the simplest machine generating your trail of bits. Meaning that at some point, they predict something distinct from your trail of bits, meaning that at some point step 2. of the process above discards them. In particular, the step after discarding \(T_{j-1}\) is arriving at \(T_j\), and staying there forever.
- Therefore, this process arrives at the simplest TM which describes the process generating your trail of bits in finite time.

QED. 

### Using the above scheme to arrive at a probability 

Now, the problem with the above scheme is that if we use our set of Turing machines to output a probability for the next bit

\[ P\Big(b_{m + 1} = 1 | (B_0, ..., B_m) \Big) := \frac{1}{n} \cdot \sum_0^n \Big(T_i(m+1) = 1\Big)  \]

then our probabilities are going to be very janky. For example, at step \( j - 1 \), that scheme would output a 0% probability to the correct value of the next bit.

To fix this, in step 2, we can require that there be not only one Turing machine that predicts all past bits, but multiple of them. What number? Well, however many your compute and memory budgets allow. This would make your probabilities less janky. 

Interestingly, that scheme also suggests that there is a tradeoff between arriving at the correct hypothesis as fast as possible—in which case we would just implement the first scheme at full speed—and producing accurate probabilities—in which case it seems like we would have to use a tweak

### A distracting epicycle: dealing with Turing machines that take too long or do not halt.

When thinking about Turing machines, one might consider one particular model, e.g., valid C programs. But in that case, it is easy to consider programs that do not halt, like:[^2]

```
void main(){
  while(1){
    ;
  }
}
```

But then in step 2 of our scheme, we never get to see what bits this program outputs, because it never outputs any bits.

This can easily be fixed as follows:

1. Start with a finite set of live Turing machines, \(\{T_0, ..., T_n\}\), a compute budget \( s \) and an empty cache of programs which take too long \( C =\{\} \)
2. Run each \( T_i \) for \( s \) seconds. 
3. If none of them predict your trail bits, \( (B_0, ..., B_m) \), within the given compute budget: 
   -  eliminate the ones who make incorrect predictions.
   -  keep track of programs which didn't output anything in the cache.
   -  attempt to compute the first \( m \) steps of Turing machine \(T_{n+1} \) with \( s \) seconds of compute. If it makes correct predictions, keep it in the set of live machines, otherwise move it to the cache.
   -  increase the compute budget to \( s + 1 \) and run each machine in the cache for one additional second
   -  Repeat step 2 until you have one program which has predicted past bits within your compute budget. Eventually this program must exist, since the Turing machine which is producing your trail of bits is by construction computable and non-halting[^3].
4. Observe the next bit, purge the machines from your set which don't predict it. If none predict it, go to 2.

It is not clear to me whether all Turing machine models require this epicycle.


<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>



[^1]: Here we assume that we have an ordering of Turing machines, i.e., that T_i is simpler than T_(i+1)
[^2]: Readers might find it amusing to run gcc loop.c and check it
[^3]: Or at least, it hasn't halted before producing the number of bits that you have seen so far.}
