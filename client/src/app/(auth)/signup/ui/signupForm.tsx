"use client";

import { useState } from "react";
import { signUp, userSignupInfo } from "../api/signUp";

export default function SignupForm() {
	const [userInfo, setUserInfo] = useState<userSignupInfo>({
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

	const handleSignUp = async (
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
			await signUp(userInfo);
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
				<div className="text-2xl">SignUp</div>

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

				<button onClick={handleSignUp} className="border-2 p-2 border-black">
					SignUp
				</button>
			</form>
		</div>
	);
}
