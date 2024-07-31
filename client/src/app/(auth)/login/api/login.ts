export interface userLoginCreds {
	username: string;
	password: string;
}

export async function login(userInfo: userLoginCreds) {
	const url = gitLoginURL();

	const query = {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(userInfo),
	};

	try {
		const response = await fetch(url, query);

		if (!response.ok) {
			const errorJson = await response.json();
			const errorMsg = errorJson?.message;

			throw new Error(errorMsg);
		}

		storeAuthToken(response);
	} catch (error) {
		throw error;
	}
}

function storeAuthToken(response: Response) {
	const authHeader = response.headers.get("Authorization");

	if (authHeader && authHeader.startsWith("Bearer ")) {
		const token = authHeader.substring(7);
		localStorage.setItem("authToken", token);
	} else {
		throw new Error("Authorization token missing in response");
	}
}

function gitLoginURL() {
	return process.env.NEXT_PUBLIC_API_URI + "/api/auth/login" || "/";
}
