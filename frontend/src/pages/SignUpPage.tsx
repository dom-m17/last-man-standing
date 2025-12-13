import { useForm, Controller } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import PhoneInput from "react-phone-number-input";
import "react-phone-number-input/style.css";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { enGB } from "date-fns/locale";
import { gql, TypedDocumentNode } from "@apollo/client";
import { useMutation } from "@apollo/client/react";

// Form validation patterns - centralized for reusability
const VALIDATION_PATTERNS = {
  name: {
    value: /^(?=.{1,30}$)[\p{L}]+(?:[ '-][\p{L}]+)*$/u,
    message: "Please enter a valid name",
  },
  username: {
    value: /^(?=.{3,30}$)[A-Za-z][A-Za-z0-9._-]*[A-Za-z0-9]$/,
    message: "Invalid username format",
  },
  password: {
    value:
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]).{8,}$/,
    message:
      "Password must be at least 8 characters and include uppercase, lowercase, number, and special character",
  },
  email: {
    value: /^\S+@\S+$/i,
    message: "Email is not in the correct format.",
  },
} as const;

// Form data types
interface SignUpFormData {
  firstName: string;
  lastName: string;
  username: string;
  password: string;
  confirmPassword: string;
  email: string;
  phoneNumber: string;
  dateOfBirth: Date | null;
  sex: string;
}

// GraphQL types
interface CreateUserVariables {
  firstName: string;
  lastName: string;
  username: string;
  password: string;
  email: string;
  phoneNumber: string;
  dateOfBirth: string;
}

interface CreateUserResponse {
  createUser: {
    id: string;
    username: string;
    firstName: string;
    lastName: string;
    email: string;
    phoneNumber: string;
    dateOfBirth: string;
    createdAt: string;
    updatedAt: string;
  };
}

const CREATE_USER: TypedDocumentNode<
  CreateUserResponse,
  CreateUserVariables
> = gql`
  mutation CreateUser(
    $firstName: String!
    $lastName: String!
    $username: String!
    $password: String!
    $email: String!
    $phoneNumber: String!
    $dateOfBirth: String!
  ) {
    createUser(
      input: {
        username: $username
        password: $password
        firstName: $firstName
        lastName: $lastName
        email: $email
        phoneNumber: $phoneNumber
        dateOfBirth: $dateOfBirth
      }
    ) {
      id
      username
      firstName
      lastName
      email
      phoneNumber
      dateOfBirth
      createdAt
      updatedAt
    }
  }
`;

export default function SignUpPage() {
  const {
    register,
    handleSubmit,
    control,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<SignUpFormData>({
    defaultValues: {
      firstName: "",
      lastName: "",
      username: "",
      password: "",
      confirmPassword: "",
      email: "",
      phoneNumber: "",
      dateOfBirth: null,
      sex: "",
    },
  });

  const [createUserMutation, { data, error, loading }] =
    useMutation(CREATE_USER);

  const password = watch("password");

  const onSubmit = async (formData: SignUpFormData) => {
    if (!formData.dateOfBirth) {
      console.error("Date of birth is required");
      return;
    }

    try {
      await createUserMutation({
        variables: {
          firstName: formData.firstName,
          lastName: formData.lastName,
          username: formData.username,
          password: formData.password,
          email: formData.email,
          phoneNumber: formData.phoneNumber,
          dateOfBirth: formData.dateOfBirth.toISOString().split("T")[0],
        },
      });

      // Handle success - redirect, show message, etc.
      console.log("User created successfully:", data);
    } catch (err) {
      // Error is already in the `error` variable from useMutation
      console.error("Failed to create user:", err);
    }
  };

  return (
    <div className="max-w-md mx-auto p-6">
      <h1 className="text-2xl font-bold mb-6">Sign Up</h1>

      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
        {/* First Name */}
        <div>
          <input
            {...register("firstName", {
              required: "First name is required",
              pattern: VALIDATION_PATTERNS.name,
            })}
            placeholder="First name"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="firstName"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Last Name */}
        <div>
          <input
            {...register("lastName", {
              required: "Last name is required",
              pattern: VALIDATION_PATTERNS.name,
            })}
            placeholder="Last name"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="lastName"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Username */}
        <div>
          <input
            {...register("username", {
              required: "Username is required",
              pattern: VALIDATION_PATTERNS.username,
            })}
            placeholder="Username"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="username"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Password */}
        <div>
          <input
            {...register("password", {
              required: "Password is required",
              pattern: VALIDATION_PATTERNS.password,
            })}
            placeholder="Password"
            type="password"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="password"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Confirm Password */}
        <div>
          <input
            {...register("confirmPassword", {
              required: "Please confirm your password",
              validate: (value) =>
                value === password || "Passwords do not match",
            })}
            placeholder="Confirm password"
            type="password"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="confirmPassword"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Email */}
        <div>
          <input
            {...register("email", {
              required: "Email is required",
              pattern: VALIDATION_PATTERNS.email,
            })}
            placeholder="Email"
            type="email"
            className="w-full px-3 py-2 border rounded"
          />
          <ErrorMessage
            errors={errors}
            name="email"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Phone Number */}
        <div>
          <Controller
            control={control}
            name="phoneNumber"
            rules={{ required: "Phone number is required" }}
            render={({ field: { onChange, value } }) => (
              <PhoneInput
                className="phone-input"
                value={value}
                onChange={onChange}
                defaultCountry="GB"
                placeholder="Enter phone number"
              />
            )}
          />
          <ErrorMessage
            errors={errors}
            name="phoneNumber"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Date of Birth */}
        <div>
          <LocalizationProvider
            dateAdapter={AdapterDateFns}
            adapterLocale={enGB}
          >
            <Controller
              name="dateOfBirth"
              control={control}
              rules={{ required: "Date of birth is required" }}
              render={({ field }) => (
                <DatePicker
                  label="Date of birth"
                  value={field.value}
                  onChange={(date) => field.onChange(date)}
                  slotProps={{
                    textField: {
                      fullWidth: true,
                    },
                  }}
                />
              )}
            />
          </LocalizationProvider>
          <ErrorMessage
            errors={errors}
            name="dateOfBirth"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Sex */}
        <div>
          <div className="flex gap-4">
            <label htmlFor="field-male" className="flex items-center gap-2">
              <input
                {...register("sex", { required: "Please select your sex" })}
                type="radio"
                value="male"
                id="field-male"
              />
              Male
            </label>
            <label htmlFor="field-female" className="flex items-center gap-2">
              <input
                {...register("sex", { required: "Please select your sex" })}
                type="radio"
                value="female"
                id="field-female"
              />
              Female
            </label>
          </div>
          <ErrorMessage
            errors={errors}
            name="sex"
            render={({ message }) => (
              <p className="text-red-500 text-sm mt-1">{message}</p>
            )}
          />
        </div>

        {/* Error display */}
        {error && (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
            <p className="font-bold">Error creating account</p>
            <p className="text-sm">{error.message}</p>
          </div>
        )}

        {/* Success display */}
        {data && (
          <div className="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded">
            <p className="font-bold">Account created successfully!</p>
            <p className="text-sm">Welcome, {data.createUser.username}!</p>
          </div>
        )}

        {/* Submit Button */}
        <button
          type="submit"
          disabled={isSubmitting || loading}
          className="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
        >
          {isSubmitting || loading ? "Creating Account..." : "Sign Up"}
        </button>
      </form>
    </div>
  );
}
