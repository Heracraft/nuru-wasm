<script>
	import { onMount } from 'svelte';

	let { code /*output=$bindable("")*/ } = $props();
	let output = $state('');

	let editorWrapper;

	let editor;

	onMount(() => {
		const go = new Go();
		WebAssembly.instantiateStreaming(fetch('/main.wasm'), go.importObject).then((result) => {
			go.run(result.instance);

			let output = runCode(code);
			console.log(output);
			
		});
		$effect(() => {
			if (code && window.runCode) {
				output = runCode(code);
				console.log(output);
				
			}
		});
	});

</script>

{output}
