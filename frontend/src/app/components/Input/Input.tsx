import styles from "./Input.module.scss";

type InputProps = {
  id?: string;
  label?: string;
  type?: "text" | "email" | "password" | "number" | "date" | "datetime-local";
  name?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
  autoComplete?: string;
  disabled?: boolean;
  placeholder?: string;
  error?: string;
};

const Input = ({
  id,
  label,
  type = "text",
  name,
  value,
  onChange,
  required,
  autoComplete,
  disabled,
  placeholder,
  error,
}: InputProps) => {
  const inputId = id ?? name ?? `input-${type}`;

  return (
    <label className={styles.label}>
      {label && (
        <span className={styles.labelText}>
          {label}
          {required && <span className={styles.required}> *</span>}
        </span>
      )}
      <input
        id={inputId}
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        required={required}
        autoComplete={autoComplete}
        disabled={disabled}
        placeholder={placeholder}
        className={styles.input}
        aria-invalid={!!error}
        aria-describedby={error ? `${inputId}-error` : undefined}
      />
      {error && (
        <span id={`${inputId}-error`} className={styles.errorText} role="alert">
          {error}
        </span>
      )}
    </label>
  );
};

export default Input;
