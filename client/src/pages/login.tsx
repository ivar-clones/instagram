import { LoginForm } from "@/components/login-form";
import { SignupRequest } from "@/models/signup-request";
import { useLogin } from "@/service/auth";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export function LoginPage() {
  const navigate = useNavigate();

  useEffect(() => {
    localStorage.getItem("token") &&
      localStorage.getItem("username") &&
      navigate("/home");
  }, []);

  const { isLoading, error, mutate } = useLogin();
  const handleLogin = (formData: SignupRequest) => {
    mutate(formData);
  };
  return (
    <div className="flex h-screen w-full items-center justify-center px-4">
      <LoginForm
        onSubmit={handleLogin}
        isLoading={isLoading}
        errorMessage={error as string}
      />
    </div>
  );
}
