import styles from "./Form.module.scss";

type FormProps = {
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
  children: React.ReactNode;
  className?: string;
};

const Form = ({ onSubmit, children, className }: FormProps) => {
  return (
    <form
      onSubmit={onSubmit}
      className={`${styles.form} ${className ?? ""}`.trim()}
    >
      {children}
    </form>
  );
};

export default Form;
