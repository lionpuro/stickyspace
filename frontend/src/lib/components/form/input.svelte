<script lang="ts">
	import type { HTMLInputAttributes } from "svelte/elements";

	type Props = {
		className?: string;
	} & HTMLInputAttributes;
	let { className = "", ...props }: Props = $props();
</script>

<input
	{...props}
	class="border-base-200 mb-6 rounded-md border p-1.5 px-2 -outline-offset-2 outline-blue-500 focus:outline-2 {className}"
/>
