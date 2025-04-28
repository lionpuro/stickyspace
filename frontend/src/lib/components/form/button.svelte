<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = {
		className?: string;
	} & HTMLButtonAttributes;
	let { className = '', children, ...props }: Props = $props();
</script>

<button
	{...props}
	class="bg-base-950 text-base-white mb-4 rounded-md px-2 py-1.75 font-medium {className}"
>
	{@render children?.()}
</button>
