<script>
	import { onMount } from 'svelte';

	import { basicSetup, EditorView } from 'codemirror';
	import { EditorState } from '@codemirror/state';
	import { javascript } from '@codemirror/lang-javascript';
	// import { vscodeDark } from '@uiw/codemirror-theme-vscode';
	import { oneDark } from '@codemirror/theme-one-dark';

	import * as Resizable from '$lib/components/ui/resizable/index.js';
	import Interpreter from '$lib/components/interpreter.svelte';

	import { File } from 'lucide-svelte';

	let code = $state('andika(5)');
	// let output = $state('');

	let editorWrapper;
	let editor;

	onMount(() => {
		let zincTheme = EditorView.theme(
			{
				'&': {
					fontFamily: 'Fira Code, monospace',
					fontSize: '1rem',
					backgroundColor: 'hsl(240 5.9% 10%) !important' // zinc-900
				},
				'.cm-content': {
					fontSize: '1.25rem'
				},
				'.cm-gutters': {
					backgroundColor: 'hsl(240 3.7% 15.9%) !important' //zinc-800
				},
				'.cm-activeLine': {
					backgroundColor: 'hsl(240 3.7% 15.9% / 50%) !important' // zinc-800
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

<Resizable.PaneGroup direction="horizontal" class="text-foreground w-full flex-1">
	<Resizable.Pane defaultSize={50}>
		<div class="mb-1 flex border-b">
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
		<div class="flex border-b">
			<div
				class="relative flex w-28 items-center justify-center gap-2 py-2.5 text-sm after:absolute after:inset-x-0 after:top-full after:h-[1px] after:w-full after:bg-yellow-600 after:content-['']"
			>
				Result
			</div>
		</div>
		<div class="flex h-full items-center justify-center p-6">
			<!-- <Interpreter {code} bind:output={output}/> -->
			<Interpreter {code} />
			<!-- <span class="font-semibold">{output}</span> -->
		</div>
	</Resizable.Pane>
</Resizable.PaneGroup>
