// frontend/app/scanner/page.tsx
import { Suspense } from "react";
import ScannerClientPage from "./scannerClient";

export default function ScannerPage() {
  return (
    <Suspense
      fallback={
        <div className="min-h-screen flex items-center justify-center text-white">
          Loading scanner...
        </div>
      }
    >
      <ScannerClientPage />
    </Suspense>
  );
}
