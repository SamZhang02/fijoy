import { createContext, useContext } from "react";
import { User } from "@/types/user";
import axios from "axios";
import { env } from "./env";
import { useQuery } from "@tanstack/react-query";

export interface AuthContext {
  user: User | undefined;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContext | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  // const [user, setUser] = useState<User | undefined>(undefined);
  // const [isLoading, setIsLoading] = useState<boolean>(true);
  //
  // useEffect(() => {
  //   axios
  //     .get(env.VITE_BACKEND_URL + "/user", {
  //       withCredentials: true,
  //     })
  //     .then((result) => {
  //       setUser(User.parse(result.data));
  //       setIsLoading(false);
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  // }, []);

  const { data: user, isLoading } = useQuery({
    queryKey: ["user"],
    queryFn: async () => {
      const result = await axios.get(env.VITE_BACKEND_URL + "/user", {
        withCredentials: true,
      });
      return User.parse(result.data);
    },
    retry: false,
  });

  return (
    <AuthContext.Provider value={{ user, isLoading }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useUser() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
