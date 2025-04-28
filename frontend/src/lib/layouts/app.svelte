<script lang="ts">
	import { session } from '$lib/stores/session.svelte';

	let { children } = $props();
</script>

<div class="flex min-h-full flex-col">
	<header class="flex px-4 py-2">
		<h1 class="flex text-xl font-semibold tracking-tight">
			Sticky
			<span class="text-base-500 ml-0.25">Space</span>
		</h1>
		<button onclick={session.logout} class="ml-auto">Log out</button>
	</header>
	{@render children?.()}
</div>
