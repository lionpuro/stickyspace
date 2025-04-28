<script lang="ts">
	import "../app.css";
	import { AppLayout, AuthLayout } from "$lib/layouts";
	import { session } from "$lib/stores/session.svelte";
	import { goto } from "$app/navigation";
	import { page } from "$app/state";
	let { children } = $props();
	let unauth = $derived(
		page.url.pathname === "/" && !session.loading && !session.user,
	);

	$effect(() => {
		if (unauth) {
			goto("/signin");
		}
	});
</script>

{#if !session.isLoggedIn}
	<AuthLayout>
		{@render children()}
	</AuthLayout>
{:else}
	<AppLayout>
		{@render children()}
	</AppLayout>
{/if}
