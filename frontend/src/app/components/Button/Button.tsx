import styles from "./Button.module.scss";

type ButtonVariant = "primary" | "secondary" | "danger";

type ButtonProps = {
  type?: "button" | "submit" | "reset";
  variant?: ButtonVariant;
  children: React.ReactNode;
  disabled?: boolean;
  className?: string;
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
};

const Button = ({
  type = "button",
  variant = "primary",
  children,
  disabled,
  className,
  onClick,
}: ButtonProps) => {
  return (
    <button
      type={type}
      disabled={disabled}
      className={`${styles.button} ${styles[variant]} ${className ?? ""}`.trim()}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button;
