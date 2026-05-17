import styles from './EmptyState.module.scss';

type EmptyStateProps = {
  message: string;
  action?: React.ReactNode;
};

const EmptyState = ({ message, action }: EmptyStateProps) => {
  return (
    <div className={styles.empty}>
      <p className={styles.message}>{message}</p>
      {action && <div className={styles.action}>{action}</div>}
    </div>
  );
};

export default EmptyState;
