<script>
	import '../styles/app.css';

	import '$lib/wasm_exec.js';
	let { children } = $props();
</script>

<div class="flex h-[100dvh] flex-col bg-zinc-900 dark">
	<div class="flex justify-between bg-zinc-800 text-foreground px-10">
		<div class="flex items-center p-4">
			<a href="/" class="text-xl flex gap-1 uppercase font-bold">
				<p class="text-yellow-500">Nuru</p>
				Playground</a>
		</div>
		<div class="flex items-center p-4 gap-10 hover:[&>a]:text-foreground text-muted-foreground">
			<a href="https://nuruprogramming.org/" target="_blank" class="text-lg">Docs</a>
			<a href="https://github.com/Heracraft/nuru-playground" target="_blank" class="text-lg">GitHub</a>
		</div>
	</div>
	{@render children()}
</div>
