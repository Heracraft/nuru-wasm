<script>
	import { onMount } from 'svelte';
	import { MediaQuery } from 'svelte/reactivity';

	import { basicSetup, EditorView } from 'codemirror';
	import { EditorState } from '@codemirror/state';
	import { javascript } from '@codemirror/lang-javascript';
	// import { vscodeDark } from '@uiw/codemirror-theme-vscode';
	import { oneDark } from '@codemirror/theme-one-dark';

	import * as Resizable from '$lib/components/ui/resizable/index.js';
	import Interpreter from '$lib/components/interpreter.svelte';
	import { Button } from '$lib/components/ui/button';

	import { File, Play } from 'lucide-svelte';

	import { defaultCode } from '$lib';

	let code = $state(defaultCode);

	let isMedium = $state(true);

	let output = $state('');

	let editorWrapper;
	let editor;

	onMount(() => {
		isMedium = new MediaQuery('min-width: 640px', true);

		// CodeMirror setup
		let zincTheme = EditorView.theme(
			{
				'&': {
					fontFamily: 'Fira Code, monospace',
					fontSize: '1rem',
					backgroundColor: 'hsl(240 5.9% 10%) !important' // zinc-900
				},
				'.cm-content': {
					// fontSize: '1.25rem'
				},

				'.cm-gutters': {
					backgroundColor: 'hsl(240 3.7% 15.9%) !important' //zinc-800
				},
				'.cm-activeLine': {
					backgroundColor: 'hsl(240 3.7% 15.9% / 50%) !important' // zinc-800
				},
				'.cm-scroller': {
					scrollbarWidth: 'thin',
					scrollbarColor: 'hsl(240 3.7% 15.9%) hsl(240 5.9% 10%)' // <accent,background> zinc-800 ,zinc-900
				}
			},
			{ dark: true }
		);

		const initialState = EditorState.create({
			doc: code,
			extensions: [
				basicSetup,
				javascript(),
				oneDark,
				zincTheme,
				EditorView.updateListener.of((v) => {
					if (v.docChanged) {
						code = v.state.doc.toString();
						// console.log({ code });
					}
				})
			]
		});
		editor = new EditorView({
			state: initialState,
			parent: editorWrapper
		});
	});
</script>

<Resizable.PaneGroup
	direction={isMedium.current ? 'horizontal' : 'vertical'}
	class="text-foreground w-full flex-1"
>
	<Resizable.Pane defaultSize={50} class="overflow-auto">
		<div class="my-1 shrink-0 border-b">
			<div
				class="relative flex w-28 items-center justify-center gap-2 py-2.5 text-sm after:absolute after:inset-x-0 after:top-full after:h-[1px] after:w-full after:bg-yellow-600 after:content-['']"
			>
				<File size={14} class="text-muted-foreground" />
				Main.nr
			</div>
		</div>
		<!-- The editor -->
		<div bind:this={editorWrapper} class="h-full w-full"></div>
	</Resizable.Pane>
	<Resizable.Handle />
	<Resizable.Pane defaultSize={50}>
		<div class="mt-1 flex justify-between border-b pb-1">
			<div
				class="relative flex w-28 items-center justify-center gap-2 py-2.5 text-sm after:absolute after:inset-x-0 after:top-full after:h-[1px] after:w-full after:bg-yellow-600 after:content-['']"
			>
				Result
			</div>
			<button
				onclick={() => {
					// if the wasm binary has loaded, then the runCode function should be available globally
					if (window.runCode) {
						
						// Dont need to update state since the program's output will get caught. See interpreter.svelte:32
						output = runCode(code);
						// console.log("Clicked: running code");
						// console.log(runCode(code));
					}
				}}
				class="hover:bg-accent mr-2 rounded p-2"
			>
				<Play size={20} class="text-yellow-500" />
			</button>
		</div>
		<!-- <div class="h-full w-full"> -->
		<Interpreter {code} {output} />
		<!-- <span class="font-semibold">{output}</span> -->
		<!-- </div> -->
	</Resizable.Pane>
</Resizable.PaneGroup>
