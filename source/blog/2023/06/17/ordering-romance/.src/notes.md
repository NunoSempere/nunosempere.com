

---

### Some details

The choice by person \(i\) for \(b\) over \(a\) could be explicit, e.g., if there was an explicit moment in which the decision was made. Or it could be implicit, e.g., if person \(i\) could have chosen to go after \(a\) but chose to go after \(b\) instead. You could also have levels of indirection, e.g. "if a prediction market was formed on whether \(i\) would choose \(a\) or \(b\) if the opportunity presented itself, would option \(b\) be above 50%?". 

The advantage of the explicit version is that it's well defined: people sometimes choose someone over someone else. The problem with it is that people don't do that often, so that set of relationships will be very sparse. In contrast, the advantage of the implicit version is that it's less sparse, but it comes at the cost of a fair amount of hand-waving, and it's not clear that we should care about it, given that it represent choices that people wouldn't naturally make.


---

equal if there is some loop
So you can't have even a partial ordering, when considering the pseudo-ordering determined by people's *actual* choices.

That said, we could try to consider the pseudo-ordering determined by people's *hypothetical* choices, e.g., as would be predicted by a prediction market. In particular, consider 

$$ a \le b \text{ iff } \sum_{i=0}^n \delta_i(a, b) < \sum_{i=0}^n \delta_i(b, a)  $$
$$ \delta_i(a, b) = \text{ 1 if person i would choose a over b, 0 otherwise }$$

Then we *still* don't get transitivity, since we can't necessarily assume that people decide rationally, i.e.,, according to VNM axioms, or even without having cyclical preferences. 

