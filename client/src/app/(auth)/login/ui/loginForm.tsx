"use client";

import { useState } from "react";
import { login, userLoginCreds as userLoginInfo } from "../api/login";

export default function LoginForm() {
	const [userInfo, setUserInfo] = useState<userLoginInfo>({
		username: "",
		password: "",
	});
	const [errorMessage, setErrorMessage] = useState("");

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;

		setUserInfo((prevInfo) => ({
			...prevInfo,
			[name]: value,
		}));
	};

	const handleLogin = async (
		e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
	) => {
		e.preventDefault();

		if (userInfo.username === "") {
			showErrorMessage("username is empty");
			return;
		}

		if (userInfo.password === "") {
			showErrorMessage("password is empty");
			return;
		}

		try {
			await login(userInfo);
		} catch (error) {
			showErrorMessage((error as Error).message);
		}
	};

	const showErrorMessage = (msg: string) => {
		setErrorMessage(msg);
	};

	return (
		<div>
			<form className="flex flex-col items-center">
				<div className="text-2xl">Login</div>

				<div className="mt-10"></div>

				<label htmlFor="username">Username:</label>
				<input
					name="username"
					onChange={handleChange}
					type="text"
					className="border-black border-2"
					value={userInfo.username}
				/>

				<label htmlFor="password">Password:</label>
				<input
					name="password"
					onChange={handleChange}
					type="password"
					className="border-black border-2"
					value={userInfo.password}
				/>

				<div className="pt-5 pb-5">{errorMessage}</div>

				<button onClick={handleLogin} className="border-2 p-2 border-black">
					Login
				</button>
			</form>
		</div>
	);
}
