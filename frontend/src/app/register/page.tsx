"use client";
import { useForm } from "react-hook-form";
import styles from "../../styles/auth.module.css";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Register() {
  const { 
    register, 
    handleSubmit, 
    formState: { errors } 
  } = useForm();
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const onSubmit = async (data: any) => {
    setIsLoading(true);
    setError("");
    
    try {
      const response = await axios.post(
        "https://friendly-doodle-pj9w6r5wxrj7h7vq9-8080.app.github.dev",
        data,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      
      if (response.status === 201) {
        router.push("/login");
      }
    } catch (err: any) {
      console.error("Erro no registro:", err);
      setError(
        err.response?.data?.message || 
        "Erro ao registrar. Verifique sua conexão e tente novamente."
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className={styles.formContainer}>
      <h2>Registro</h2>
      {error && <div className={styles.error}>{error}</div>}
      
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          {...register("email", { required: "Email é obrigatório" })}
          type="email"
          placeholder="Email"
          className={styles.inputField}
        />
        {errors.email && (
          <span className={styles.error}>
            {String(errors.email.message)}
          </span>
        )}

        <input
          {...register("password", { 
            required: "Senha é obrigatória",
            minLength: {
              value: 6,
              message: "Senha deve ter pelo menos 6 caracteres"
            }
          })}
          type="password"
          placeholder="Senha"
          className={styles.inputField}
        />
        {errors.password && (
          <span className={styles.error}>
            {String(errors.password.message)}
          </span>
        )}

        <button 
          type="submit" 
          className={styles.button}
          disabled={isLoading}
        >
          {isLoading ? "Registrando..." : "Registrar"}
        </button>
      </form>
    </div>
  );
}