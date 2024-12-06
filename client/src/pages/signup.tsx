import { SignupForm } from "@/components/signup-form";
import { SignupRequest } from "@/models/signup-request";
import { useSignup } from "@/service/auth";

export function SignupPage() {
  const { isLoading, error, mutate } = useSignup();

  const handleSignup = (formData: SignupRequest) => {
    mutate(formData);
  };

  return (
    <div className="flex h-screen w-full items-center justify-center px-4">
      <SignupForm
        onSubmit={handleSignup}
        isLoading={isLoading}
        errorMessage={error as string}
      />
    </div>
  );
}
