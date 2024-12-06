import { Link } from "react-router";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { SignupRequest } from "@/models/signup-request";
import { useState } from "react";
import { Loader2 } from "lucide-react";

interface SignupFormProps {
  readonly onSubmit: (data: SignupRequest) => void;
  readonly isLoading: boolean;
  readonly errorMessage: string;
}

export function SignupForm(props: SignupFormProps) {
  const { onSubmit, isLoading, errorMessage } = props;
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [retypePasswordError, setRetypePasswordError] = useState("");

  const handleSignup = () => {
    if (password.length > 0 && password !== confirmPassword) {
      setRetypePasswordError("Passwords don't match");
      return;
    }

    if (username && password && password === confirmPassword) {
      onSubmit({ username, password });
    }
  };

  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Signup</CardTitle>
        <CardDescription>
          Enter username and password to create your account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid gap-4">
          <div className="grid gap-2">
            <Label htmlFor="username">Username</Label>
            <Input
              id="username"
              type="username"
              required
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            {errorMessage && <p className="text-red-500">{errorMessage}</p>}
          </div>
          <div className="grid gap-2">
            <div className="flex items-center">
              <Label htmlFor="password">Password</Label>
            </div>
            <Input
              id="password"
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="grid gap-2">
            <div className="flex items-center">
              <Label htmlFor="confrmPassword">Re-enter password</Label>
            </div>
            <Input
              id="confrmPassword"
              type="password"
              required
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            {retypePasswordError && (
              <p className="text-red-500">{retypePasswordError}</p>
            )}
          </div>
          <Button type="submit" className="w-full" onClick={handleSignup}>
            {isLoading && <Loader2 className="animate-spin" />}
            Signup
          </Button>
        </div>
        <div className="mt-4 text-center text-sm">
          Already have an account?{" "}
          <Link to="/login" className="underline">
            Login
          </Link>
        </div>
      </CardContent>
    </Card>
  );
}
