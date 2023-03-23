Find a beta distribution that fits your desired confidence interval
===================================================================

Here is a tool for finding a beta distribution that fits your desired confidence interval. 
E.g., to find a beta distribution whose 95% confidence interval is 0.2 to 0.8, 
input 0.2, 0.8, and 0.95 in their respective fields below:

<form action="/signup" method="post" id="fit-beta"> 
  <div class="field">
    <label for="ci_lower">Value at the lower end of your confidence interval</label>
    <input type="text" id="ci_lower" ci_lower="ci_lower" placeholder="0.2 to ..." class="subscribe-input"/>
  </div>
  <br>
  <div class="field">
    <label for="ci_upper">Value at the upper end of your confidence interval</label>
    <input type="text" id="ci_upper" ci_upper="ci_upper" placeholder="... to 0.8" class="subscribe-input"/>
  </div>
  <br>
  <div class="field">
    <label for="ci_length">Length of your confidence interval (0.9 = 90% by default)</label>
    <input type="text" id="ci_length" ci_length="ci_length" placeholder="0.9" class="subscribe-input"/>
  </div>
  <br>
  <button type="submit" id="submit_button"class="subscribe-button">Calculate</button>
</form>
<p id="result"></p>

<script>

const form = document.getElementById('fit-beta');
const submit_button = document.getElementById('submit_button')
const result_p = document.getElementById('result')

form.addEventListener('submit', (event) => {
	event.preventDefault();
	
	result_p.innerHTML = ''
	// Disable new submits
	submit_button.disabled = 'disabled';
	// console.log(submit_button.innerHTML)
	submit_button.innerHTML = 'Calculating...'

	// get the form data
	let data = {
		ci_lower: form.elements['ci_lower'].value,
		ci_upper: form.elements['ci_upper'].value,
		ci_length: form.elements['ci_length'].value,
	}
	console.log(JSON.stringify(data))

		// make request
  fetch("https://trastos.nunosempere.com/fit-beta", {
		method: "post",
		headers: {
			'Accept': 'application/json',
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data)
	})
	.then( (response) => response.json())
	.then(result => {
		submit_button.disabled = false;
		submit_button.innerHTML = 'Calculate'
		result_p.innerHTML = `Result: beta(${result[0]}, ${result[1]})`
		console.log(result)
		console.log(JSON.stringify(result, null, 2))
	});
});
</script>

<p>The virtue of using a a beta distribution is that it is inherently <em>bounded</em> between 0 and 1,
and thus much more suitable for estimating things like probabilities, and, more speculatively,
ranges of values. Previously, various people in my extended social circle, including myself,  had been
using lognormals (bad, because not bounded) or truncated lognormals (better, but deeply inelegant) for
those purposes. I hope that with this tool such practices will come to and end.</p>

You can read more about how this utility works [here](https://github.com/quantified-uncertainty/fit-beta.git).
Advanced users will particularly enjoy the [npm package](https://www.npmjs.com/package/fit-beta).

<p><img src="https://i.imgur.com/3uQgbcw.png" style="width: 20%;">
<br>
This is a project of the <a href="https://quantifieduncertainty.org/">Quantified Uncertainty Research Institute</a>.</p>

<p>
  <section id='isso-thread'>
  <noscript>Javascript needs to be activated to view comments.</noscript>
  </section>
</p>

