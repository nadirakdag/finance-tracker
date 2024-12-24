import React, { useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { TransactionFormData, TransactionType } from '../../types/transaction';
import LoadingSpinner from '../common/LoadingSpinner';

interface TransactionFormProps {
  onSubmit: (type: TransactionType, data: TransactionFormData) => Promise<void>;
  isLoading?: boolean;
}

const TransactionForm: React.FC<TransactionFormProps> = ({ onSubmit, isLoading }) => {
  const [transaction, setTransaction] = useState<TransactionFormData>({
    amount: '',
    description: '',
    category: '',
    date: new Date().toISOString().split('T')[0]
  });

  const handleSubmit = async (type: TransactionType) => {
    try {
      await onSubmit(type, transaction);
      setTransaction({
        amount: '',
        description: '',
        category: '',
        date: new Date().toISOString().split('T')[0]
      });
    } catch (error) {
      console.error('Failed to submit transaction:', error);
    }
  };

  if (isLoading) {
    return (
      <Card>
        <CardContent className="p-6">
          <LoadingSpinner />
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Add Transaction</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <Input
            type="number"
            placeholder="Amount"
            value={transaction.amount}
            onChange={(e) => setTransaction({...transaction, amount: e.target.value})}
          />
          <Input
            type="text"
            placeholder="Description"
            value={transaction.description}
            onChange={(e) => setTransaction({...transaction, description: e.target.value})}
          />
          <Select
            value={transaction.category}
            onValueChange={(value) => setTransaction({...transaction, category: value})}
          >
            <SelectTrigger>
              <SelectValue placeholder="Category" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="food">Food</SelectItem>
              <SelectItem value="transport">Transport</SelectItem>
              <SelectItem value="utilities">Utilities</SelectItem>
              <SelectItem value="entertainment">Entertainment</SelectItem>
            </SelectContent>
          </Select>
          <Input
            type="date"
            value={transaction.date}
            onChange={(e) => setTransaction({...transaction, date: e.target.value})}
          />
          <div className="flex space-x-2">
            <Button onClick={() => handleSubmit('income')} className="bg-green-600 hover:bg-green-700">
              Add Income
            </Button>
            <Button onClick={() => handleSubmit('expense')} className="bg-red-600 hover:bg-red-700">
              Add Expense
            </Button>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default TransactionForm;