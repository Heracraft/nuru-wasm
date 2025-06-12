<script>
	import { onMount } from 'svelte';

	import { fly } from 'svelte/transition';

	import { Progress } from '$lib/components/ui/progress/index.js';
	import { Info } from 'lucide-svelte';

	let { code, output } = $props();
	// let output = $state('');
	let loadProgress = $state(0);

	let editorWrapper;

	let editor;

	function setUp() {
		// runs before the wasm binary is initialized
		// It's role is to register the output (nuruOutputReceiver) capture function

		window.nuruOutputReceiver = function (output) {
			console.log("outpit",output);
			
			output += `\n${output}`;
		};
	}

	async function loadWasmBinary(url) {
		// Start the fetch request
		const response = await fetch(url);

		if (!response.ok) {
			throw new Error(`Failed to fetch ${url}`);
		}

		// Get the total size from the headers
		const contentLength = response.headers.get('Content-Length');
		if (!contentLength) {
			throw new Error('Content-Length header is missing');
		}

		// Create a reader to track the stream
		const reader = response.body.getReader();
		const total = parseInt(contentLength, 10);
		let received = 0;

		// Create an array buffer to hold the WASM bytes
		let chunks = [];

		// Read the stream in chunks
		while (true) {
			const { done, value } = await reader.read();
			if (done) break;

			// Track received bytes
			chunks.push(value);
			received += value.length;

			// Update progress bar
			const percentComplete = Math.round((received / total) * 100);
			loadProgress = percentComplete;
		}

		// Combine the chunks into a single buffer
		const wasmBytes = new Uint8Array(received);
		let position = 0;
		for (let chunk of chunks) {
			wasmBytes.set(chunk, position);
			position += chunk.length;
		}

		loadProgress = 100; // fallback

		return wasmBytes;
	}

	onMount(async () => {
		const go = new Go();

		// `.instantiateStreaming` way more efficient at fetching but no way to track progress gotta use fetch/XHR
		// WebAssembly.instantiateStreaming(fetch('/main.wasm'), go.importObject).then((result) => {
		// 	go.run(result.instance);
		// 	runCode(code);
		// });

		// using fetch
		const wasmBytes = await loadWasmBinary('/main.wasm');

		setUp()

		WebAssembly.instantiate(wasmBytes.buffer, go.importObject).then((result) => {
			go.run(result.instance);
			// output=runCode(code);
		});

		// Auto-run code. But is too annoying especially when you have user prompts.

		// $effect(() => {
		// 	if (code && window.runCode) {
		// 		setTimeout(() => {
		// 			runCode(code);
		// 		}, 1000);
		// 	}
		// });

		// Intercepting console.logs from "wasm_exec.js" - don't do this kids
		// Temp solution removed

		// const log = console.log.bind(console);
		// console.log = (...args) => {
		// 	log(...args);
		// 	output = args.join(' ');
		// };

		async () => {
			(await fetch('/main.wasm')).headers.get('Content-Length');
		};
	});
</script>

<div class="relative h-full w-full p-2">
	{#if loadProgress != 100}
		<div out:fly={{ y: -5 }} class="absolute inset-x-0 top-0 flex flex-col gap-2 bg-accent p-2">
			<div class="flex items-center gap-2">
				<Info size={16} />
				<p>Loading the interpreter - {loadProgress}%</p>
			</div>
			<Progress value={loadProgress} class="h-2"></Progress>
		</div>
	{/if}
	{output}
</div>
