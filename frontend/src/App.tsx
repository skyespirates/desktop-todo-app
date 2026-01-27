import { Button } from "@/components/ui/button";
import List from "@/components/List.tsx";
import { Input } from "@/components/ui/input";
import { useEffect, useState } from "react";
import { v4 as uuid } from "uuid";
import type { KeyboardEvent } from "react";
import { repository } from "../wailsjs/go/models";
import { CreateTodo, ListTodos } from "../wailsjs/go/main/App";
import Modal from "@/components/Modal";

export type Todo = {
  id: string;
  title: string;
  description: string;
  completed?: boolean;
  createdAt?: Date;
  completedAt?: Date;
};

const App = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [title, setTitle] = useState("");
  const [open, setOpen] = useState(false);
  const [selectedTodo, setSelectedTodo] = useState<Todo>({} as Todo);

  useState;
  useEffect(() => {
    ListTodos()
      .then((list) => setTodos(list))
      .then((err) => console.log(err));
  }, []);

  async function handleAddTodo() {
    try {
      const cleanTitle = title.trim();
      if (!cleanTitle) {
        return;
      }
      const todo = {
        id: uuid(),
        title: cleanTitle,
        description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.",
      };

      const td = new repository.Todo(todo);

      await CreateTodo(td);
      setTitle("");
      ListTodos()
        .then((list) => setTodos(list))
        .then((err) => console.log(err));
    } catch (error) {
      console.log(error);
    }
  }

  // async function fetchTodo() {
  //   try {
  //     const td = await ListTodos();
  //     console.log("todos", td);
  //   } catch (error) {
  //     console.log(error);
  //   }
  // }

  function handleKeyDown(e: KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      handleAddTodo();
    }
  }
  return (
    <div className="flex flex-col justify-center items-center gap-4">
      <h1 className="text-3xl text-center">Todo Apps</h1>
      <div className="flex gap-2 min-w-96">
        <Input
          type="text"
          placeholder="so, what are you gonna do today? 🐔"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          onKeyDown={handleKeyDown}
        />
        <Button className="cursor-pointer" onClick={handleAddTodo}>
          Add Todo
        </Button>

        {/* <Button className="cursor-pointer" onClick={fetchTodo}>
          Fetch Todos
        </Button> */}
      </div>
      <div>
        <Modal open={open} setOpen={setOpen} todo={selectedTodo} />
      </div>
      {todos && (
        <List
          items={todos}
          open={open}
          setOpen={setOpen}
          setTodos={setTodos}
          listTodos={ListTodos}
          setSelectedTodo={setSelectedTodo}
        />
      )}
    </div>
  );
};

export default App;
