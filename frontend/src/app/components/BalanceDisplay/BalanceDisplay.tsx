import { formatYen } from '@/app/libs/format';
import styles from './BalanceDisplay.module.scss';

type BalanceDisplayProps = {
  label: string;
  amount: number;
};

const BalanceDisplay = ({ label, amount }: BalanceDisplayProps) => {
  return (
    <div className={styles.wrapper}>
      <p className={styles.label}>{label}</p>
      <p className={styles.amount}>{formatYen(amount)}</p>
    </div>
  );
};

export default BalanceDisplay;
