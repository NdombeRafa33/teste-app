"use client";  // ✅ Isso força o componente a rodar no lado do cliente
import { useForm } from "react-hook-form";
import styles from "../../styles/auth.module.css";
import axios from "axios";
import { useRouter } from "next/navigation";


export default function Login() {
  const { register, handleSubmit } = useForm();
  const router = useRouter();

  const onSubmit = async (data: any) => {
    try {
      const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, data);
      const token = response.data.token;
      localStorage.setItem("token", token);
      router.push("/dashboard");
    } catch (error) {
      alert("Erro ao fazer login, tente novamente.");
    }
  };

  return (
    <div className={styles.formContainer}>
      <h2>Login</h2>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          {...register("email")}
          type="email"
          placeholder="Email"
          className={styles.inputField}
        />
        <input
          {...register("password")}
          type="password"
          placeholder="Senha"
          className={styles.inputField}
        />
        <button type="submit" className={styles.button}>
          Entrar
        </button>
      </form>
    </div>
  );
}
