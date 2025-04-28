<script lang="ts">
	import { goto } from '$app/navigation';
	import { Form, FormButton, FormInput } from '$lib/components';
	import { auth } from '$lib/firebase-client';
	import { IconGoogle } from '$lib/icons';
	import { session } from '$lib/stores/session.svelte';
	import {
		GoogleAuthProvider,
		createUserWithEmailAndPassword,
		signInWithPopup,
	} from 'firebase/auth';

	async function signupWithEmail(
		e: SubmitEvent & { currentTarget: HTMLFormElement },
	) {
		e.preventDefault();
		const email = e.currentTarget['email'].value;
		const password = e.currentTarget['password'].value;
		if (!email || !password) {
			return;
		}

		try {
			const creds = await createUserWithEmailAndPassword(auth, email, password);
			session.user = creds.user;
			goto('/');
		} catch (err) {
			console.error(err);
			return err;
		}
	}

	async function signinWithGoogle() {
		const provider = new GoogleAuthProvider();
		try {
			const creds = await signInWithPopup(auth, provider);
			session.user = creds.user;
			goto('/');
		} catch (err) {
			console.error(err);
			return err;
		}
	}
</script>

<div class="m-auto flex w-full max-w-md flex-col">
	<div
		class="bg-base-white flex w-full max-w-md flex-col rounded-lg p-6 sm:p-8"
	>
		<h1 class="mb-6 text-xl font-bold">Sign up</h1>
		<Form onsubmit={signupWithEmail}>
			<label for="email" class="mb-2">Email</label>
			<FormInput id="email" name="email" type="email" />
			<label for="password" class="mb-2">Password</label>
			<FormInput id="password" name="password" type="password" />
			<FormButton>Sign up</FormButton>
		</Form>
		<span
			class="
			text-base-500 before:bg-base-200 after:bg-base-200 mb-4 flex items-center justify-between gap-2
			text-sm font-medium before:h-px before:grow
			before:content-['_'] after:h-px after:grow after:content-['_']
		"
		>
			or
		</span>
		<button
			onclick={signinWithGoogle}
			class="border-base-200 text-base-600 flex justify-center gap-2 rounded-md border px-2 py-1.5 font-medium"
		>
			<IconGoogle />
			Sign in with Google
		</button>
	</div>
	<span class="text-base-600 mt-2 flex justify-center gap-2">
		Already have an account?
		<a href="/signin" class="text-primary-500 font-medium">Sign in</a>
	</span>
</div>
