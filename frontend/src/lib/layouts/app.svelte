<script lang="ts">
	import { session } from "$lib/stores/session.svelte";

	let { children } = $props();
</script>

<div class="flex min-h-full flex-col">
	<header class="flex px-4 py-2">
		<h1 class="text-xl font-semibold tracking-tight">
			<a href="/" class="flex">
				<span class="text-base-900">sticky</span>
				<span class="text-base-500 ml-0.25">space</span>
			</a>
		</h1>
		<button onclick={session.logout} class="ml-auto">Log out</button>
	</header>
	{@render children?.()}
</div>
