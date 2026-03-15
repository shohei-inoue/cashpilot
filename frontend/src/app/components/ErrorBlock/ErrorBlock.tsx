import styles from "./ErrorBlock.module.scss";

type ErrorBlockProps = {
  children: React.ReactNode;
};

const ErrorBlock = ({ children }: ErrorBlockProps) => {
  return (
    <p className={styles.errorBlock} role="alert">
      {children}
    </p>
  );
};

export default ErrorBlock;
