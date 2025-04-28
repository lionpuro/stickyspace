<script lang="ts">
	let { children } = $props();
</script>

<div class="flex min-h-full flex-col p-4">
	{@render children?.()}
</div>
