"use client";

import * as React from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  // DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

type ModalProps = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
};

const Modal = ({ open, setOpen }: ModalProps) => {
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      {/* <DialogTrigger asChild>
        <Button variant="outline">Edit Todo</Button>
      </DialogTrigger> */}
      <DialogContent className="sm:max-w-106.25">
        <DialogHeader>
          <DialogTitle>Edit Todo</DialogTitle>
          <DialogDescription>
            Make changes to your todo here. Click save when you&apos;re done.
          </DialogDescription>
        </DialogHeader>
        <EditForm />
      </DialogContent>
    </Dialog>
  );
};

const EditForm = () => {
  return (
    <div className="grid items-start gap-6">
      <div className="grid gap-3">
        <Label htmlFor="title">Title</Label>
        <Input id="title" defaultValue="walking on the beach" />
      </div>
      <div className="grid gap-3">
        <Label htmlFor="description">Description</Label>
        <Input
          id="description"
          defaultValue="just chilling and enjoying life"
        />
      </div>
      <Button type="button">Save changes</Button>
    </div>
  );
};

export default Modal;
