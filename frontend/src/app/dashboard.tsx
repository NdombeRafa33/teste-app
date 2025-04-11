import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import styles from "../styles/dashboard.module.css";

export default function Dashboard() {
  const router = useRouter();
  const [userData, setUserData] = useState<any>(null);

  useEffect(() => {
    // Verificar se o token existe
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login"); // Se não houver token, redireciona para login
    } else {
      // Aqui você pode fazer uma requisição para a API para pegar os dados do usuário
      setUserData({ email: "user@example.com" }); // Exemplo de dados do usuário
    }
  }, [router]);

  return (
    <div className={styles.dashboardContainer}>
      <div className={styles.dashboardHeader}>
        <h1>Bem-vindo ao seu Dashboard!</h1>
      </div>
      <div className={styles.dashboardContent}>
        {userData ? (
          <div>
            <p>Você está logado como: {userData.email}</p>
            {/* Aqui você pode adicionar mais funcionalidades do dashboard */}
            <a href="/profile" className={styles.dashboardButton}>
              Editar Perfil
            </a>
          </div>
        ) : (
          <p>Carregando dados...</p>
        )}
      </div>
    </div>
  );
}
