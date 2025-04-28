<script lang="ts">
	import type { HTMLFormAttributes } from 'svelte/elements';

	let { children, onsubmit, ...props }: HTMLFormAttributes = $props();
</script>

<form {...props} {onsubmit} class="contents">
	{@render children?.()}
</form>
