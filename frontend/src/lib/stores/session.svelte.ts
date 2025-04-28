import { onAuthStateChanged, signOut, type User } from "firebase/auth";
import { auth } from "$lib/firebase-client";

class AuthSession {
	user: User | null = $state(null);
	loading: boolean = $state(true);
	isLoggedIn: boolean = $state(
		(() => {
			const val = localStorage.getItem("authenticated");
			return val !== null && JSON.parse(val) === true;
		})(),
	);

	constructor() {
		onAuthStateChanged(auth, (user) => {
			this.user = user;
			this.setLoggedIn(user !== null);
			this.loading = false;
		});
	}

	setLoggedIn(val: boolean) {
		this.isLoggedIn = val;
		localStorage.setItem("authenticated", JSON.stringify(val));
	}

	async logout() {
		await signOut(auth);
	}
}

export const session = new AuthSession();
