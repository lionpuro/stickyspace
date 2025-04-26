<script lang="ts">
	import { AppLayout } from '$lib/layouts';
	import '../app.css';

	let { children } = $props();
</script>

<AppLayout>
	{@render children()}
</AppLayout>
