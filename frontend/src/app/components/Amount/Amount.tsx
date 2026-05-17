import { formatYen } from '@/app/libs/format';
import styles from './Amount.module.scss';

type AmountVariant = 'income' | 'expense' | 'neutral';

type AmountProps = {
  amount: number;
  variant?: AmountVariant;
  className?: string;
};

const Amount = ({ amount, variant = 'neutral', className }: AmountProps) => {
  const resolvedVariant =
    variant === 'neutral'
      ? amount > 0
        ? 'income'
        : amount < 0
          ? 'expense'
          : 'neutral'
      : variant;

  return (
    <span
      className={`${styles.amount} ${styles[resolvedVariant]} ${className ?? ''}`.trim()}
    >
      {formatYen(amount)}
    </span>
  );
};

export default Amount;
