import Amount from '@/app/components/Amount/Amount';
import { formatDateOnly } from '@/app/libs/format';
import type { GoalProjection } from '@/app/types/simulation';
import styles from './GoalProjectionList.module.scss';

type GoalProjectionListProps = {
  projections: GoalProjection[];
};

const GoalProjectionList = ({ projections }: GoalProjectionListProps) => {
  if (projections.length === 0) {
    return (
      <p className={styles.empty}>
        期限付きの目標が設定画面にないため、達成見込みは表示されません。
      </p>
    );
  }

  return (
    <ul className={styles.list}>
      {projections.map((goal) => (
        <li key={goal.goal_id} className={styles.item}>
          <div className={styles.main}>
            <span className={styles.name}>{goal.name}</span>
            <span className={styles.meta}>
              目標 <Amount amount={goal.target_amount} variant="income" />
              {goal.deadline && (
                <span className={styles.deadline}>
                  期限 {formatDateOnly(goal.deadline)}
                </span>
              )}
            </span>
            <span className={styles.projected}>
              予測残高 <Amount amount={goal.projected_balance} />
            </span>
          </div>
          <span
            className={`${styles.badge} ${goal.achievable ? styles.achievable : styles.unachievable}`}
          >
            {goal.achievable ? '達成見込み' : '未達の見込み'}
          </span>
        </li>
      ))}
    </ul>
  );
};

export default GoalProjectionList;
