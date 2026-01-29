"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Checkbox } from "@/components/ui/checkbox";
import { Field } from "@/components/ui/field";
import { type Todo } from "@/App";
import { UpdateTodo, ListTodos } from "../../wailsjs/go/main/App";

type ModalProps = {
  open: boolean;
  todo: Todo;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setTodos: React.Dispatch<React.SetStateAction<Todo[]>>;
};

const Modal = ({ open, todo, setOpen, setTodos }: ModalProps) => {
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description);
  const [completed, setCompleted] = useState(todo.completed);

  useEffect(() => {
    if (!todo) return;
    setTitle(todo.title);
    setDescription(todo.description);
    setCompleted(todo.completed);
  }, [todo]);

  function handleClick() {
    const data = {
      id: todo.id,
      title,
      description,
      completed,
    };

    UpdateTodo(data)
      .then((res) => {
        console.log("Success", res);
        setOpen((prev) => !prev);
        ListTodos().then((todos) => setTodos(todos));
      })
      .catch((err) => console.log("Error", err));
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent
        className="sm:max-w-106.25"
        onOpenAutoFocus={(e) => e.preventDefault()}
      >
        <DialogHeader>
          <DialogTitle>Edit Todo</DialogTitle>
          <DialogDescription>
            Make changes to your todo here. Click save when you&apos;re done.
          </DialogDescription>
        </DialogHeader>
        <div className="grid items-start gap-6">
          <div className="grid gap-3">
            <Label htmlFor="title">Title</Label>
            <Input
              id="title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
          </div>
          <div className="grid gap-3">
            <Label htmlFor="description">Description</Label>
            <Textarea
              placeholder="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </div>
          <div className="grid gap-3">
            <Field orientation="horizontal">
              <Checkbox
                checked={completed}
                className="cursor-pointer"
                id="completed"
                onClick={() => {
                  setCompleted((prev) => !prev);
                }}
              />
              <Label htmlFor="completed">Completed</Label>
            </Field>
          </div>
          <Button
            onClick={handleClick}
            type="button"
            className="cursor-pointer"
          >
            Save changes
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default Modal;
