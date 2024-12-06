import { useQuery } from "react-query";
import { useNavigate } from "react-router";

export const usePosts = () => {
  const navigate = useNavigate();

  return useQuery({
    queryKey: ["posts"],
    queryFn: async () => {
      const response = await fetch("http://localhost:8080/api/v1/posts", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });

      if (!response.ok) {
        if (response.status === 401) {
          localStorage.removeItem("token");
          localStorage.removeItem("username");
          return navigate("/login");
        }
        return Promise.reject("something went wrong");
      }

      return response.json();
    },
  });
};
