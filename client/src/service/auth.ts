import { SignupRequest } from "@/models/signup-request";
import { useMutation } from "react-query";
import { useNavigate } from "react-router";

export const useSignup = () => {
  const navigate = useNavigate();

  return useMutation((data: SignupRequest) => {
    return fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((data) => Promise.reject(data.error));
        }

        return response.json();
      })
      .then((data) => {
        localStorage.setItem("token", data.token);
        localStorage.setItem("username", data.username);
        navigate("/home");
      });
  });
};

export const useLogin = () => {
  const navigate = useNavigate();

  return useMutation((data: SignupRequest) => {
    return fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((data) => Promise.reject(data.error));
        }

        return response.json();
      })
      .then((data) => {
        localStorage.setItem("token", data.token);
        localStorage.setItem("username", data.username);
        navigate("/home");
      });
  });
};
