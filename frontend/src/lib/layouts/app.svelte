<script lang="ts">
	let { children } = $props();
</script>

<div class="flex min-h-full flex-col">
	<header class="border-base-100 flex border-b px-4 py-2">
		<h1 class="flex text-xl font-semibold tracking-tight">
			Sticky
			<span class="text-base-500 ml-0.25">Space</span>
		</h1>
	</header>
	{@render children?.()}
</div>
