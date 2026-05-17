import { formatYen } from '@/app/libs/format';
import type { CashflowTrendPoint } from '@/app/actions/dashboard';
import styles from './CashflowChart.module.scss';

type CashflowChartProps = {
  points: CashflowTrendPoint[];
};

const CashflowChart = ({ points }: CashflowChartProps) => {
  if (points.length === 0) {
    return (
      <p className={styles.empty}>取引データが増えると、残高推移が表示されます。</p>
    );
  }

  const values = points.map((p) => p.balance);
  const min = Math.min(...values, 0);
  const max = Math.max(...values, 0);
  const range = max - min || 1;

  return (
    <div className={styles.chart}>
      <div className={styles.bars} role="img" aria-label="月次残高推移">
        {points.map((point) => {
          const heightPct = ((point.balance - min) / range) * 100;
          return (
            <div key={point.period} className={styles.barGroup}>
              <div className={styles.barTrack}>
                <div
                  className={styles.bar}
                  style={{ height: `${Math.max(heightPct, 4)}%` }}
                  title={`${point.label}: ${formatYen(point.balance)}`}
                />
              </div>
              <span className={styles.label}>
                {(() => {
                  const [, m] = point.period.split('-');
                  return m ? `${Number(m)}月` : point.label;
                })()}
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default CashflowChart;
