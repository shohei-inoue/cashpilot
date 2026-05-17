import styles from './Select.module.scss';

export type SelectOption = {
  value: string;
  label: string;
};

type SelectProps = {
  id?: string;
  label?: string;
  name?: string;
  value: string;
  options: readonly SelectOption[];
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  required?: boolean;
  disabled?: boolean;
  error?: string;
};

const Select = ({
  id,
  label,
  name,
  value,
  options,
  onChange,
  required,
  disabled,
  error,
}: SelectProps) => {
  const selectId = id ?? name ?? 'select';

  return (
    <label className={styles.label}>
      {label && (
        <span className={styles.labelText}>
          {label}
          {required && <span className={styles.required}> *</span>}
        </span>
      )}
      <select
        id={selectId}
        name={name}
        value={value}
        onChange={onChange}
        required={required}
        disabled={disabled}
        className={styles.select}
        aria-invalid={!!error}
        aria-describedby={error ? `${selectId}-error` : undefined}
      >
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
      {error && (
        <span id={`${selectId}-error`} className={styles.errorText} role="alert">
          {error}
        </span>
      )}
    </label>
  );
};

export default Select;
